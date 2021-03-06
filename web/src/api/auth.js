import fetch from "./fetch";

export const apiLogin = data => {
	return fetch.post("/auth/login", data);
};
export const admAuth = () => {
	return fetch.get("/auth/login");
};
export const apiLogoff = data => {
	//   return fetch.post("/api/logoff", data);
};
// 统计状态
export const admCollect = () => {
	return fetch.get("/adm/collect");
};
// 服务器信息
export const admSys = () => {
	return fetch.get("/system");
};
