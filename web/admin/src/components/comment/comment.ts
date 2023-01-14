import day from 'dayjs';

const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '6%',
    align: 'center',
  },
  {
    title: '创建时间',
    dataIndex: 'CreatedAt',
    width: '12%',
    align: 'center',
    customRender: (data:any) => {
      return data.value ? day(data.value).format('YYYY年MM月DD日 HH:mm') : '暂无'
    }
  },
  {
    title: '评论文章',
    dataIndex: 'article_title',
    width: '15%',
    align: 'center',
  },
  {
    title: '评论者',
    dataIndex: 'user_username',
    width: '10%',
    align: 'center',
  },
  {
    title: '评论内容',
    dataIndex: 'content',
    width: '20%',
    align: 'center',
  },
  {
    title: '评论状态',
    dataIndex: 'status',
    width: '10%',
    align: 'center',
  },
  {
    title: '操作',
    width: '25%',
    align: 'center',
    dataIndex: 'operation',
    scopedSlots: {
      customRender: 'operation'
    },
    fixed: 'right',
  },
]

export { columns }