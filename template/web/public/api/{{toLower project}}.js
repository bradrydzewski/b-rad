import { instance } from "./config.js";
import useSWR, { mutate } from "swr";

/**
 * create{{title project}} creates a new {{toLower project}}.
 */
export const create{{title project}} = async (data, fetcher) => {
	return fetcher(`${instance}/api/v1/{{toLower project}}s`, {
		body: JSON.stringify(data),
		method: "POST",
	}).then(({{toLower project}}) => {
		mutate(`${instance}/api/v1/user/{{toLower project}}s`);
		return {{toLower project}};
	});
};

/**
 * update{{title project}} updates an existing {{toLower project}}
 */
export const update{{title project}} = (params, data, fetcher) => {
	const { id } = params;
	return fetcher(`${instance}/api/v1/{{toLower project}}s/${id}`, {
		body: JSON.stringify(data),
		method: "PATCH",
	}).then((_) => {
		mutate(`${instance}/api/v1/user/{{toLower project}}s`);
		mutate(`${instance}/api/v1/user/{{toLower project}}s/${id}`);
		return;
	});
};

/**
 * delete{{title project}} deletes an existing {{toLower project}}
 */
export const delete{{title project}} = (params, fetcher) => {
	const { id } = params;
	return fetcher(`${instance}/api/v1/{{toLower project}}s/${id}`, {
		method: "DELETE",
	}).then((_) => {
		mutate(`${instance}/api/v1/user/{{toLower project}}s`);
		return;
	});
};

/**
 * use{{title project}}List returns an swr hook that provides a {{toLower project}} list.
 */
export const use{{title project}}List = () => {
	const { data, error } = useSWR(`${instance}/api/v1/user/{{toLower project}}s`);
	return {
		{{toLower project}}List: data,
		isLoading: !error && !data,
		isError: error,
	};
};

/**
 * use{{title project}} returns an swr hook that provides the {{toLower project}}.
 */
export const use{{title project}} = (id) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/{{toLower project}}s/${id}?token=true`
	);
	return {
		{{toLower project}}: data,
		isLoading: !error && !data,
		isError: error,
	};
};
