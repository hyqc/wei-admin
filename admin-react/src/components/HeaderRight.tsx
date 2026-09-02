import { useNavigate } from 'react-router-dom';
import { Avatar, Dropdown, message } from 'antd';
import type { MenuProps } from 'antd';
import { UserOutlined, LogoutOutlined } from '@ant-design/icons';
import { useUserStore } from '@/store/user';
import { LoginPath } from '@/api/config';

/** 头部右侧：账号下拉（个人中心 / 退出登录） */
export default function HeaderRight() {
  const navigate = useNavigate();
  const userInfo = useUserStore((s) => s.userInfo);
  const logoutAsync = useUserStore((s) => s.logoutAsync);

  const items: MenuProps['items'] = [
    { key: 'account', icon: <UserOutlined />, label: '个人中心' },
    { type: 'divider' },
    { key: 'logout', icon: <LogoutOutlined />, label: '退出登录' },
  ];

  const onClick: MenuProps['onClick'] = async ({ key }) => {
    if (key === 'account') {
      navigate('/account');
      return;
    }
    if (key === 'logout') {
      await logoutAsync();
      message.success('退出登录成功');
      navigate(LoginPath);
    }
  };

  const displayName = userInfo?.nickname || userInfo?.username;

  return (
    <div className="header-right">
      <Dropdown menu={{ items, onClick }} placement="bottomRight">
        <div className="account">
          <Avatar src={userInfo?.avatar} size={28} className="avatar">
            {(displayName || '?')[0]}
          </Avatar>
          <span className="name">{displayName}</span>
        </div>
      </Dropdown>
    </div>
  );
}
