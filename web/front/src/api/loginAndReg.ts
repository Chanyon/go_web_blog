import { request } from "./request";

type LoginAndRegResponse = {
  data: {
    data: string,
    id: number,
    status: number,
    message: string,
  }
};

type RegisterResponse = {
  data: {
    status: number,
    message: string,
  }
};

export const loginApi = (loginForm:object):Promise<LoginAndRegResponse> => {
  return request({
    url:"/user/login",
    method:"POST",
    data: {
      ...loginForm
    }
  });
};

export const registerApi = (registerForm:object): Promise<RegisterResponse> => {
  return request({
    url:"/user/add",
    method:"POST",
    data: {
      ...registerForm
    }
  });
}