import day from 'dayjs';

type UserInfo = {
  CreatedAt?: string;
  DeletedAt?: null;
  id?: number;
  UpdatedAt?: string;
  password: string;
  checkpass?: string;
  role?: number;
  username: string;
}

const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '10%',
    align: 'center',
  },
  {
    title: '用户名',
    dataIndex: 'username',
    width: '20%',
    align: 'center',
  },
  {
    title: '注册时间',
    dataIndex: 'CreatedAt',
    width: '20%',
    align: 'center',
    customRender: (data:any) => {
      return data.value ? day(data.value).format('YYYY年MM月DD日 HH:mm') : '暂无'
    }
  },
  {
    title: '角色',
    dataIndex: 'role',
    width: '10%',
    align: 'center',
  },
  {
    title: '操作',
    width: '30%',
    align: 'center',
    dataIndex: 'operation',
    scopedSlots: {
      customRender: 'operation'
    },
    fixed: 'right',
  },
]
export {columns}
export type {UserInfo}