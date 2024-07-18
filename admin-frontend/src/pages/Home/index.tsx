import React, { useState } from 'react';
import { Container, Content } from '@/components/PageListContainer';
import { Alert, Calendar } from 'antd';
import { Moment } from 'moment';
import moment from 'moment';

const Demo: React.FC = () => {
  const [value, setValue] = useState(moment(new Date()));
  const [selectedValue, setSelectedValue] = useState(moment(new Date()));

  const onSelect = (newValue: Moment) => {
    setValue(newValue);
    setSelectedValue(newValue);
  };

  const onPanelChange = (newValue: Moment) => {
    setValue(newValue);
  };
  return (
    <Container>
      <Content>
        <Alert message={`当前日期: ${selectedValue?.format('YYYY-MM-DD')}`} />
        <Calendar fullscreen value={value} onSelect={onSelect} onPanelChange={onPanelChange} />
      </Content>
    </Container>
  );
};

export default Demo;
