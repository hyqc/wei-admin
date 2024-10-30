import React, {  } from 'react';
import { Container, Content } from '@/components/PageListContainer';
import { Calendar, CalendarProps } from 'antd';
import { Dayjs } from 'dayjs';

const Demo: React.FC = () => {

  const onPanelChange = (value: Dayjs, mode: CalendarProps<Dayjs>['mode']) => {
    console.log(value.format('YYYY-MM-DD'), mode);
  };

  return (
    <Container>
      <Content>
        <Calendar fullscreen onPanelChange={onPanelChange} /> 
      </Content> 
    </Container>
  );
};

export default Demo;
