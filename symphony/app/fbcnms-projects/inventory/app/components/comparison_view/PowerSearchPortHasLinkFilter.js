/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type {FilterProps} from '../comparison_view/ComparisonViewTypes';

import * as React from 'react';
import PowerSearchBoolFilter from './PowerSearchBoolFilter';

const PowerSearchPortHasLinkFilter = (props: FilterProps) => {
  return <PowerSearchBoolFilter label="Is Linked" {...props} />;
};

export default PowerSearchPortHasLinkFilter;
