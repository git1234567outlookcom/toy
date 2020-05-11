import fetch from "./fetch";
// 分类信息
export const apiCateAll = () => {
	return fetch.get("/category/list");
};
// 添加分类
export const admCateAdd = data => {
	return fetch.post(`/category`, data);
};
// 修改分类
export const admCateEdit = data => {
	return fetch.post(`/category`, data);
};
// 删除分类
export const admCateDrop = id => {
	return fetch.delete(`/category/${id}`);
};
