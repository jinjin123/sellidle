import React from 'react';
import store from './store';
import { observer } from 'mobx-react';
import { Col, Button, Row, Tag } from 'antd';


@observer
class PortTable extends React.Component {
  componentDidMount() {
    store.fetchRecords()
  }
  render() {
    let data = store.records;
    return (
      <div >
        <Button type="primary" onClick={() => store.showBindPortForm()}>映射端口</Button>
        <Row justify="start" >
          {
            data.map((item, index) => (
              <Col span={2} key={index}>
                <span>{item.Port}
                  {item.State === "close" ? <Tag color="volcano">{"close"}</Tag> : <Tag color="green">{"open"}</Tag>}
                </span>
              </Col>
            ))
          }
        </Row>
      </div>
    );
  }

}

export default PortTable