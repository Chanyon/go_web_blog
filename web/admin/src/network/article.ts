import {request} from './request'

export const getCategoryInfo = (route:string) => {
  return request({
    url: route,
    method: 'get',
  });
}

export const getArticleByPage = (route:string,data:object) => {
  return request({
    url: route,
    method: 'get',
    params:{
      ...data,
    }
  });
}

export const getArticleByCid = (route:string,data:object) => {
  return request({
    url: route,
    method: 'get',
    params:{
      ...data,
    }
  });
}

export const deleteArticleById = (route:string) => {
  return request({
    url: route,
    method: 'delete',
  });
}
export const getArticleInfoById = (route:string) => {
  return request({
    url: route,
    method: 'get',
  });
}

export const uploadArticleById = (route:string,data:object) => {
  return request({
    url: route,
    method: 'put',
    data:{
      ...data,
    }
  });
}

export const addArticleInfo = (route:string,data:object) => {
  return request({
    url: route,
    method: 'post',
    data:{
      ...data,
    }
  });
}