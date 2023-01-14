import {request} from './request'

export const Login = (route:string,data:object) => {
  return request({
    url: route,
    method: 'post',
    data:{
      ...data
    }
  });
}