import { http } from "@/utils/http";

type Result = {
  code: number;
  msg: string;
  data: Array<any>;
};

export const getAsyncRoutes = () => {
  return http.request<Result>("get", "/api/get-async-routes");
};
