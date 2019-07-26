import { instance } from "./config.js";
import useSWR, { mutate } from "swr";

/**
 * create{{parent}} creates a new {{toLower parent}}.
 */
export const create{{parent}} = async (params, data, fetcher) => {
	const { project } = params;
	return fetcher(`${instance}/api/v1/projects/${project}/{{toLower parent}}s`, {
		body: JSON.stringify(data),
		method: "POST",
	}).then(({{toLower parent}}) => {
		mutate(`${instance}/api/v1/projects/${project}/{{toLower parent}}s`);
		return {{toLower parent}};
	});
};

/**
 * update{{parent}} updates an existing {{toLower parent}}.
 */
export const update{{parent}} = (params, data, fetcher) => {
	const { project, {{toLower parent}} } = params;
	return fetcher(`${instance}/api/v1/projects/${project}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}`, {
		body: JSON.stringify(data),
		method: "PATCH",
	});
};

/**
 * delete{{parent}} deletes an existing {{toLower parent}}.
 */
export const delete{{parent}} = (params, fetcher) => {
	const { project, {{toLower parent}} } = params;
	return fetcher(`${instance}/api/v1/projects/${project}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}`, {
		method: "DELETE",
	}).then((_) => {
		mutate(`${instance}/api/v1/projects/${project}/{{toLower parent}}s`);
		return;
	});
};

/**
 * use returns an swr hook that provides
 */
export const use{{parent}}List = (project) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/projects/${project}/{{toLower parent}}s`
	);

	return {
		{{toLower parent}}List: data,
		isLoading: !error && !data,
		isError: error,
	};
};

/**
 * use returns an swr hook that provides
 */
export const use{{parent}} = (project, {{toLower parent}}) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/projects/${project}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}`
	);

	return {
		{{toLower parent}}: data,
		isLoading: !error && !data,
		isError: error,
	};
};
