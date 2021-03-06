/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

const {LOG_FORMAT, LOG_LEVEL} = require('./config');

// This must be done before any module imports to configure
// logging correctly
const logging = require('@fbcnms/logging');
logging.configure({
  LOG_FORMAT,
  LOG_LEVEL,
});

const {
  appMiddleware,
  csrfMiddleware,
  organizationMiddleware,
  sessionMiddleware,
} = require('@fbcnms/express-middleware');
const connectSession = require('connect-session-sequelize');
const express = require('express');
const passport = require('passport');
const path = require('path');
const fbcPassport = require('@fbcnms/auth/passport').default;
const session = require('express-session');
const {sequelize, Organization} = require('@fbcnms/sequelize-models');
const OrganizationLocalStrategy = require('@fbcnms/auth/strategies/OrganizationLocalStrategy')
  .default;
const OrganizationSamlStrategy = require('@fbcnms/auth/strategies/OrganizationSamlStrategy')
  .default;

const {createGraphTenant, deleteGraphTenant} = require('./graphgrpc/tenant');
const {access, configureAccess} = require('@fbcnms/auth/access');
const {
  AccessRoles: {USER},
} = require('@fbcnms/auth/roles');

const devMode = process.env.NODE_ENV !== 'production';

// Create Sequelize Store
const SessionStore = connectSession(session.Store);
const sequelizeSessionStore = new SessionStore({db: sequelize});

// add hooks to Organization model
Organization.beforeCreate((org: any) => {
  createGraphTenant(org.name);
});
Organization.beforeDestroy((org: any) => {
  deleteGraphTenant(org.name);
});

// FBC express initialization
const app = express();
app.set('trust proxy', 1);
app.use(organizationMiddleware());
app.use(appMiddleware());
app.use(
  sessionMiddleware({
    devMode,
    sessionStore: sequelizeSessionStore,
    sessionToken:
      process.env.SESSION_TOKEN || 'fhcfvugnlkkgntihvlekctunhbbdbjiu',
  }),
);
app.use(passport.initialize());
app.use(passport.session()); // must be after sessionMiddleware

fbcPassport.use();
passport.use('local', OrganizationLocalStrategy());
passport.use(
  'saml',
  OrganizationSamlStrategy({
    urlPrefix: '/user',
  }),
);

// Views
app.set('views', path.join(__dirname, '..', 'views'));
app.set('view engine', 'pug');

// Routes
app.use(
  '/inventory/static',
  express.static(path.join(__dirname, '..', 'static')),
);
app.use('/user', require('@fbcnms/auth/express').unprotectedUserRoutes());

app.use(configureAccess({loginUrl: '/user/login'}));

// All /graph endpoints don't use CORS and are JSON (no form),
// so no CSRF is needed
app.use('/graph', access(USER), require('./graph/routes'));

app.use('/', csrfMiddleware(), access(USER), require('./main/routes').default);

// Catch All
// TODO: Disabling now until we find a better way to handle paths that dont
// exist
// app.get('*', (req, res) => {
//   res.redirect('/inventory');
// });

export default app;
