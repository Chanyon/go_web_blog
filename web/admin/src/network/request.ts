import axios, { AxiosRequestConfig} from 'axios'

export function request(config:AxiosRequestConfig<any>) {
  const base:AxiosRequestConfig<any> = {
      baseURL:"http://127.0.0.1:3000/api/v1",
      timeout: 6000
  }
  const instance = axios.create(base);

  instance.interceptors.request.use(config => {
    let token = window.sessionStorage.getItem("token");
    
    config.headers!.Authorization = `Bearer ${token}`;
    return config;
  });

  instance.interceptors.response.use(res => {
    return res;
  },err => {
    return err;
  });
  return instance(config);
}