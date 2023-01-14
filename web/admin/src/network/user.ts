import {request} from './request'

export const getUserList = (route:string,data: object) => {
  return request({
    url: route,
    method: 'get',
    params:{
      ...data
    }
  });
}
export const getSearchData = (route:string,data: object) => {
  return request({
    url: route,
    method: 'get',
    params:{
      ...data
    }
  });
}

export const UpdateUserPass = (route:string,data: object) => {
  return request({
    url: route,
    method: 'put',
    data:{
      ...data
    }
  });
}

export const updateUserInfo = (route:string,data: object) => {
  return request({
    url: route,
    method: 'put',
    data:{
      ...data
    }
  });
}

export const getUserInfoById = (route:string) => {
  return request({
    url: route,
    method: 'get',
  });
}


export const deleteUserInfoById = (route:string) => {
  return request({
    url: route,
    method: 'deleted',
  });
}

export const addUserInfo = (route:string,data: object) => {
  return request({
    url: route,
    method: 'post',
    data:{
      ...data
    }
  });
}

export const getProfile = (route:string) => {
  return request({
    url: route,
    method: 'get',
  });
}

export const uploadProfileInfo = (route:string,data: object) => {
  return request({
    url: route,
    method: 'post',
    data:{
      ...data
    }
  });
}
export const uploadProfileById = (route:string,data: object) => {
  return request({
    url: route,
    method: 'put',
    data:{
      ...data
    }
  });
}