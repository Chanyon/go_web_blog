interface DataItem {
  id: string;
  name: string;
}

const columns = [
  {
    title: '分类ID',
    dataIndex: 'id',
    width: '20%',
    align: 'center',
  },
  {
    title: '分类名',
    dataIndex: 'name',
    align: 'center',
  },
  {
    title: '操作',
    dataIndex: 'operation',
    align: 'center',
  },
];

export type {DataItem}
export {
  columns
}