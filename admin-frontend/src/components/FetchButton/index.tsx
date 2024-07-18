import { Button } from 'antd';
import React from 'react';
const FetchButton: React.FC<any> = (props) => {
  const color = props.danger ? '#FF4D4F' : '#1890FF';
  const styleCss = {
    padding: '0 4px',
    margin: '0 2px',
    color,
  };
  return (
    <Button type="text" style={styleCss} {...props}>
      {props?.children}
    </Button>
  );
};

export default FetchButton;
