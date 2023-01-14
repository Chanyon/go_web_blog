import {request} from './request'

export const getComment = (route:string,data:object) => {
  return request({
    url: route,
    method: 'get',
    params:{
      ...data,
    }
  });
}

export const updateComment = (route:string,data:object) => {
  return request({
    url: route,
    method: 'put',
    data:{
      ...data,
    }
  });
}

export const deleteComment = (route:string) => {
  return request({
    url: route,
    method: 'delete',
  });
}
