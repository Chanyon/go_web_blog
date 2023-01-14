import { request } from "./request"
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

type CateResponseI = {
  data: {
    data: Array<ArticleI & CategoryI>, 
    message: string,
    status: number,
    total: number,
  }
};
export const getCategoryApi = (): Promise<CateResponseI> => {
  return request({
    method: "GET",
    url: "user/category",
  });
};