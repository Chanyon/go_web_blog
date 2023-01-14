import { request } from "./request";
export type ProfileI = {
  data: {
    data: {
      avatar: string,
      bili: string,
      desc: string,
      email: string,
      icp_record: string,
      id: string,
      img: string,
      name: string,
      qq_chat: string,
      wei_bo: string,
      wei_chat: string,
    },
    message: string,
    status: number,
  }
};

export const getProfileApi = (route: string): Promise<ProfileI> => {
  return request({
    url: route,
    method:"get",
  });
};