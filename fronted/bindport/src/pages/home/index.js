import React from 'react';
import PortTable from './Table';
import PortForm from './Form';
import store from './store';
import { observer } from 'mobx-react';

@observer
class HomeIndex extends React.Component {
  render() {
    return (
      <div>
        {store.PortForm && <PortForm />}
        <PortTable />
      </div>
    )
  }
}

export default HomeIndex
