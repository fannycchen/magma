/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */
import type {AlertConfig} from './AlarmAPIType';

import Grid from '@material-ui/core/Grid';
import PrometheusEditor from './PrometheusEditor';
import React from 'react';
import {cloneDeep} from 'lodash';
import {makeStyles} from '@material-ui/styles';
import {useEnqueueSnackbar} from '@fbcnms/ui/hooks/useSnackbar';
import {useRouter} from '@fbcnms/ui/hooks';
import {useState} from 'react';
import type {ApiUtil} from './AlarmsApi';

type Props = {
  apiUtil: ApiUtil,
  additionalAlertRuleOptions?: Object,
  onExit: () => void,
  initialConfig: ?AlertConfig,
  isNew: boolean,
  thresholdEditorEnabled: boolean,
};

const useStyles = makeStyles(theme => ({
  header: {
    padding: `${theme.spacing(2)}px ${theme.spacing(3)}px`,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between',
    backgroundColor: 'white',
    borderBottom: `1px solid ${theme.palette.divider}`,
  },
  gridContainer: {
    flexGrow: 1,
  },
  editingSpace: {
    height: '100%',
    padding: '30px',
  },
}));

export default function AddEditAlert(props: Props) {
  const {isNew, apiUtil, onExit} = props;
  const [alertConfig, setAlertConfig] = useState<?AlertConfig>(
    cloneDeep(props.initialConfig),
  );

  const {match} = useRouter();
  const enqueueSnackbar = useEnqueueSnackbar();
  const saveAlert = async () => {
    try {
      if (!alertConfig) {
        throw new Error('Alert config empty');
      }

      const request = {
        networkId: match.params.networkId,
        rule: alertConfig,
      };
      if (isNew) {
        await apiUtil.createAlertRule(request);
      } else {
        await apiUtil.editAlertRule(request);
      }
      enqueueSnackbar(`Successfully ${isNew ? 'added' : 'saved'} alert rule`, {
        variant: 'success',
      });
      onExit();
    } catch (error) {
      enqueueSnackbar(
        `Unable to create alert: ${
          error.response ? error.response.data.message : error.message
        }.`,
        {
          variant: 'error',
        },
      );
    }
  };

  const classes = useStyles();

  return (
    <Grid
      className={classes.gridContainer}
      container
      spacing={0}
      data-testid="add-edit-alert">
      <Grid className={classes.editingSpace} item xs>
        <PrometheusEditor
          apiUtil={props.apiUtil}
          onExit={props.onExit}
          updateAlertConfig={setAlertConfig}
          isNew={props.isNew}
          saveAlertRule={saveAlert}
          rule={alertConfig}
          thresholdEditorEnabled={props.thresholdEditorEnabled}
        />
      </Grid>
    </Grid>
  );
}
