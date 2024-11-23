import { Button } from "antd";

const FetchButton: React.FC<any> = (props) => {
    const color = props.danger ? '#FF4D4F' : '#1890FF';
    const styleCss = {
      padding: '1px',
      margin: '0 4px',
      fontSize: '12px',
      color,
    };
    return (
      <Button type="text" style={styleCss} {...props}>
        {props?.children}
      </Button>
    );
  };

export default FetchButton;