/**
 * Copyright 2004-present Facebook. All Rights Reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * @flow
 * @format
 */

import type PowerSearchLinksResultsTable_links from './__generated__/PowerSearchLinksResultsTable_links.graphql';
import type {AppContextType} from '@fbcnms/ui/context/AppContext';
import type {ContextRouter} from 'react-router-dom';
import type {WithAlert} from '@fbcnms/ui/components/Alert/withAlert';
import type {WithStyles} from '@material-ui/core';

import AppContext from '@fbcnms/ui/context/AppContext';
import Box from '@material-ui/core/Box';
import EquipmentBreadcrumbs from '../equipment/EquipmentBreadcrumbs';
import React from 'react';
import Text from '@fbcnms/ui/components/design-system/Text';
import classNames from 'classnames';
import withAlert from '@fbcnms/ui/components/Alert/withAlert';
import {AutoSizer, Column, Table} from 'react-virtualized';
import {capitalize} from '@fbcnms/util/strings';
import {createFragmentContainer, graphql} from 'react-relay';
import {getPropertyValue} from '../../common/Property';
import {lowerCase} from 'lodash';
import {withRouter} from 'react-router-dom';
import {withStyles} from '@material-ui/core/styles';

import 'react-virtualized/styles.css';

const styles = theme => ({
  root: {
    width: '100%',
    marginTop: theme.spacing(3),
    overflowX: 'auto',
  },
  table: {
    outline: 'none',
  },
  header: {
    borderBottom: '2px solid #f0f0f0',
  },
  cell: {
    padding: '16px 17px',
    lineHeight: '100%',
  },
  addButton: {
    paddingLeft: '16px',
    paddingRight: '16px',
  },
  futureState: {
    textTransform: 'capitalize',
    maxWidth: '50px',
  },
  icon: {
    padding: '0px',
    marginLeft: theme.spacing(),
  },
  headerText: {
    fontSize: '12px',
    lineHeight: '16px',
    color: 'rgba(0, 0, 0, 0.54)',
    textTransform: 'none',
  },
  cellText: {
    fontSize: '13px',
    lineHeight: '16px',
  },
  linkText: {
    color: 'inherit',
    display: 'inline',
    fontWeight: 'bold',
  },
  propsCell: {
    padding: '16px 17px',
    display: 'block',
    overflow: 'auto',
    lineHeight: '100%',
  },
});

type Props = WithAlert &
  WithStyles<typeof styles> &
  ContextRouter & {
    links: PowerSearchLinksResultsTable_links,
  };

class PowerSearchLinksResultsTable extends React.Component<Props> {
  static contextType = AppContext;
  context: AppContextType;

  _headerRenderer = ({label}) => {
    const {classes} = this.props;
    return (
      <div className={classes.cell}>
        <Text className={classes.headerText}>{label}</Text>
      </div>
    );
  };

  _cellRenderer = ({dataKey, rowData, cellData}) => {
    const {classes, history} = this.props;
    let content = null;

    if (dataKey === 'portType' || dataKey === 'equipmentType') {
      content = (
        <Text className={classNames(classes.cellText)}>{cellData}</Text>
      );
    } else if (dataKey === 'status' && rowData.futureState) {
      content = (
        <Text className={classNames(classes.cellText, classes.linkText)}>
          {cellData}
        </Text>
      );
    } else if (dataKey === 'equipmentA' || dataKey === 'equipmentB') {
      const index = dataKey === 'equipmentA' ? 0 : 1;
      content = (
        <EquipmentBreadcrumbs
          equipment={rowData.ports[index].parentEquipment}
          showSelfEquipment={true}
          onParentLocationClicked={locationId =>
            history.push(
              `inventory/` + (locationId ? `?location=${locationId}` : ''),
            )
          }
          onEquipmentClicked={equipmentId =>
            history.push(
              `inventory/` + (equipmentId ? `?equipment=${equipmentId}` : ''),
            )
          }
          size="small"
        />
      );
    } else if (dataKey === 'properties') {
      return (
        <div className={classes.propsCell}>
          {cellData.map(property => {
            const {name} = property.propertyType;
            const val = getPropertyValue(property) ?? '';
            return <Box>{`${name}: ${val}`}</Box>;
          })}
        </div>
      );
    } else {
      content = <Text className={classes.cellText}>{cellData}</Text>;
    }

    return <div className={classes.cell}>{content}</div>;
  };

  _getRowHeight = rowData => {
    return rowData.properties.length > 3
      ? 40 + rowData.properties.length * 10
      : 50;
  };

  render() {
    const {classes, links} = this.props;
    if (links.length === 0) {
      return null;
    }
    const equipmetStatusEnabled = this.context.isFeatureEnabled(
      'planned_equipment',
    );

    return links.length > 0 ? (
      <AutoSizer>
        {({height, width}) => (
          <Table
            className={classes.table}
            height={height}
            width={width}
            headerHeight={50}
            rowHeight={({index}) => this._getRowHeight(links[index])}
            rowCount={links.length}
            rowGetter={({index}) => links[index]}
            gridClassName={classes.table}
            rowClassName={({index}) => (index === -1 ? classes.header : '')}>
            <Column
              label="Equipment A"
              dataKey="equipmentA"
              width={350}
              flexGrow={1}
              headerRenderer={this._headerRenderer}
              cellRenderer={this._cellRenderer}
              cellDataGetter={({rowData}) =>
                rowData.ports[0].parentEquipment.name
              }
            />
            <Column
              label="Equipment A Type"
              dataKey="equipmentType"
              width={250}
              flexGrow={1}
              headerRenderer={this._headerRenderer}
              cellRenderer={this._cellRenderer}
              cellDataGetter={({rowData}) =>
                rowData.ports[0].parentEquipment.equipmentType.name
              }
            />
            <Column
              label="Port A Name"
              dataKey="portType"
              width={150}
              flexGrow={1}
              headerRenderer={this._headerRenderer}
              cellRenderer={this._cellRenderer}
              cellDataGetter={({rowData}) => rowData.ports[0].definition.name}
            />
            <Column
              label="Equipment B"
              dataKey="equipmentB"
              width={350}
              flexGrow={1}
              headerRenderer={this._headerRenderer}
              cellRenderer={this._cellRenderer}
              cellDataGetter={({rowData}) =>
                rowData.ports[1].parentEquipment.name
              }
            />
            <Column
              label="Equipment B Type"
              dataKey="equipmentType"
              width={250}
              flexGrow={1}
              headerRenderer={this._headerRenderer}
              cellRenderer={this._cellRenderer}
              cellDataGetter={({rowData}) =>
                rowData.ports[1].parentEquipment.equipmentType.name
              }
            />
            <Column
              label="Port B Name"
              dataKey="portType"
              width={150}
              flexGrow={1}
              headerRenderer={this._headerRenderer}
              cellRenderer={this._cellRenderer}
              cellDataGetter={({rowData}) => rowData.ports[1].definition.name}
            />
            <Column
              label="Properties"
              dataKey="properties"
              width={350}
              flexGrow={1}
              headerRenderer={this._headerRenderer}
              cellRenderer={this._cellRenderer}
              cellDataGetter={({rowData}) => rowData.properties}
            />
            {equipmetStatusEnabled && (
              <Column
                label="Status"
                dataKey="status"
                cellDataGetter={({rowData}) =>
                  rowData.futureState && rowData.workOrder
                    ? `${capitalize(
                        lowerCase(rowData.workOrder.status),
                      )} ${lowerCase(rowData.futureState)}`
                    : 'Installed'
                }
                width={250}
                flexGrow={1}
                headerRenderer={this._headerRenderer}
                cellRenderer={this._cellRenderer}
              />
            )}
          </Table>
        )}
      </AutoSizer>
    ) : null;
  }
}

export default withRouter(
  withAlert(
    withStyles(styles)(
      createFragmentContainer(PowerSearchLinksResultsTable, {
        links: graphql`
          fragment PowerSearchLinksResultsTable_links on Link
            @relay(plural: true) {
            id
            futureState
            ports {
              id
              definition {
                id
                name
                visibleLabel
                type
                portType {
                  linkPropertyTypes {
                    ...PropertyTypeFormField_propertyType @relay(mask: false)
                  }
                }
              }
              parentEquipment {
                id
                name
                futureState
                equipmentType {
                  id
                  name
                  portDefinitions {
                    id
                    name
                    visibleLabel
                    type
                    bandwidth
                    portType {
                      id
                      name
                    }
                  }
                }
                ...EquipmentBreadcrumbs_equipment
              }
            }
            properties {
              id
              stringValue
              intValue
              floatValue
              booleanValue
              latitudeValue
              longitudeValue
              rangeFromValue
              rangeToValue
              propertyType {
                id
                name
                type
                isEditable
                isInstanceProperty
                stringValue
                intValue
                floatValue
                booleanValue
                latitudeValue
                longitudeValue
                rangeFromValue
                rangeToValue
              }
            }
            workOrder {
              id
              status
            }
          }
        `,
      }),
    ),
  ),
);
