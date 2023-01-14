import { request } from "./request";
export type ArticleI = {
  Category: CategoryI,
  ID: number,
  cid?: number,
  comment_count: number,
  content: string,
  desc: string,
  img: string,
  read_count: number,
  title: string,
  CreatedAt?: string,
  UpdatedAt?: string,
  DeletedAt?: string,
}
export type CategoryI = {
  ID?: number,
  id: number,
  name: string,
  CreatedAt?: string,
  UpdatedAt?: string,
  DeletedAt?: string,
}
type ArticleResponseI = {
  data:{
    data: ArticleI,
    message: string,
    status: number,
  }
}
type CommentI = {
  data: {
    data : any[],
    total: number,
    message: string,
    status: number,
  }
}

export const getArticleInfoApi = (route: string):Promise<ArticleResponseI> => {
  return request({
    url: route,
    method:"GET",
    data:{
    
    },
  });
};

export const getCommentListApi = (route: string):Promise<CommentI> => {
  return request({
    url: route,
    method:"GET",
    data:{
    
    },
  });
};

export const sendCommentApi = (commData: object):Promise<CommentI> => {
  return request({
    url: "user/comment/create",
    method:"POST",
    data:{
      ...commData,
    },
  });
};
