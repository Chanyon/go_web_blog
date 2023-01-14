import day from 'dayjs';

interface CategoryItem {
  CreatedAt?:string;
  DeletedAt?:string;
  UpdatedAt?:string;
  ID?:number;
  id:number;
  name:string;
}
interface ArticleItem {
  CreatedAt?:string;
  DeletedAt?:string;
  UpdatedAt:string;
  ID:number;
  cid:number;
  title:string;
  desc:string;
  name:string;
  img:string;
}
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '100px',
    align: 'center',
    fixed: 'left'
  },
  {
    title: '更新日期',
    dataIndex: 'UpdatedAt',
    width: '10%',
    align: 'center',
    customRender: (data: any) => {
      return data.value ? day(data.value).format('YYYY年MM月DD日 HH:mm') : '暂无'
    },
  },
  {
    title: '分类',
    width: '15%',
    dataIndex: 'name',
    align: 'center',
  },
  {
    title: '文章标题',
    dataIndex: 'title',
    width: '25%',
    align: 'center',
  },
  {
    title: '文章描述',
    dataIndex: 'desc',
    width: '20%',
    align: 'center',
  },
  {
    title: '缩略图',
    dataIndex: 'img',
    align: 'center',
    width: '100px',
    scopedSlots: {
      customRender: 'img'
    }
  },
  {
    title: '操作',
    width: '15%',
    dataIndex: 'operation',
    align: 'center',
    scopedSlots: {
      customRender: 'operation'
    },
    fixed: 'right',
  },
];
export type {CategoryItem,ArticleItem}
export {
  columns
}