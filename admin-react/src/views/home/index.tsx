import { useState } from 'react';
import { Card, Calendar, Tag } from 'antd';
import dayjs, { type Dayjs } from 'dayjs';

/** 首页：排班日历 */
export default function Home() {
  const today = dayjs();
  const [currentDate, setCurrentDate] = useState<Dayjs>(dayjs());

  return (
    <div className="home-page">
      <Card title="排班日历">
        <Calendar
          fullscreen
          value={currentDate}
          onChange={(v) => setCurrentDate(v as Dayjs)}
          cellRender={(current: Dayjs, info) => {
            if (info.type === 'date' && current.date() === today.date()) {
              return (
                <ul style={{ margin: 0, padding: 0, listStyle: 'none' }}>
                  <li style={{ marginTop: 4 }}>
                    <Tag color="blue">今日无排班计划</Tag>
                  </li>
                </ul>
              );
            }
            return info.originNode;
          }}
        />
      </Card>
    </div>
  );
}
