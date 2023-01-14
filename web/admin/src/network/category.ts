import {request} from './request'

export const getCategoryInfo = (route:string,data:object) => {
  return request({
    url: route,
    method: 'get',
    params:{
      ...data
    }
  });
}

export const addCategory = (route:string,data:object) => {
  return request({
    url: route,
    method: 'post',
    data:{
      ...data
    }
  });
}

export const getCategoryById = (route:string) => {
  return request({
    url: route,
    method: 'get',
  });
}

export const editCategoryById = (route:string,data:object) => {
  return request({
    url: route,
    method: 'put',
    data:{
    ...data
    }
  });
}

export const deleteCateById = (route:string) => {
  return request({
    url: route,
    method: 'delete',
  });
}