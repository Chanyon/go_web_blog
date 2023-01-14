import { request } from "./request";
type ArticleI = {
  ID: number,
  cid: number,
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
type CategoryI = {
  ID?: number,
  id: number,
  name: string,
  CreatedAt?: string,
  UpdatedAt?: string,
  DeletedAt?: string,
}

export type ArticleResponse = {
  data: {
    data: Array<ArticleI>,
    status: number,
    message: string,
    total: number,
  }
};

export const getArticleListApi = (route:string): Promise<ArticleResponse> => {
  return request({
    url: route,
    method:"GET",
    data: {
    }
  });
}