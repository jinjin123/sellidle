/**
 * Copyright (c) OpenSpug Organization. https://github.com/openspug/spug
 * Copyright (c) <spug.dev@gmail.com>
 * Released under the AGPL-3.0 License.
 */
import React from 'react';
import { observer } from 'mobx-react';
import { InputNumber, Modal, Form, Input, message } from 'antd';
import http from 'libs/http';
import store from './store';


@observer
class PortForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: false,
    }
  }

  handleSubmit = () => {
    this.setState({ loading: true });
    const formData = this.props.form.getFieldsValue();
    if (formData["ip"] == "192.168.1.110") {
      message.error("未知错误,核查正确ip")
      return
    }
    formData["inside"] = formData["inside"].toString()
    formData["outside"] = formData["outside"].toString()
    http.post('/api/update', formData)
      .then(res => {
        if (res == 401) {
          message.error("无法绑定,检查地址和端口是否有效")
          this.setState({ loading: false });
          return
        }
        setTimeout(() => {
          message.success('绑定成功');
          store.PortForm = false;
          store.fetchRecords()
        },
          3000)
      }, () => this.setState({ loading: false }))
  };

  render() {
    const { getFieldDecorator } = this.props.form;
    return (
      <Modal
        visible
        width={800}
        maskClosable={false}
        title={'新建应用'}
        onCancel={() => store.PortForm = false}
        confirmLoading={this.state.loading}
        onOk={this.handleSubmit}>
        <Form labelCol={{ span: 6 }} wrapperCol={{ span: 14 }}>
          <Form.Item required label="外部端口">
            {getFieldDecorator('outside')(
              <InputNumber placeholder="17001-17999范围 一次只能开一个" style={{ "width": 350 }} />
            )}
          </Form.Item>
          <Form.Item required label="内部主机IP">
            {getFieldDecorator('ip')(
              <Input placeholder="ifconfig 看你机器ip 192.168.122.x" />
            )}
          </Form.Item>
          <Form.Item required label="内部主机端口">
            {getFieldDecorator('inside')(
              <InputNumber placeholder="666" style={{ "width": 350 }} />
            )}
          </Form.Item>
        </Form>
      </Modal>
    )
  }
}

export default Form.create()(PortForm)
