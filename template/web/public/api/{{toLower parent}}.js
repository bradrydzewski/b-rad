import { instance } from "./config.js";
import useSWR, { mutate } from "swr";

/**
 * create{{title parent}} creates a new {{toLower parent}}.
 */
export const create{{title parent}} = async (params, data, fetcher) => {
	const { {{toLower project}} } = params;
	return fetcher(`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s`, {
		body: JSON.stringify(data),
		method: "POST",
	}).then((response) => {
		mutate(`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s`);
		return response;
	});
};

/**
 * update{{title parent}} updates an existing {{toLower parent}}.
 */
export const update{{title parent}} = (params, data, fetcher) => {
	const { {{toLower project}}, {{toLower parent}} } = params;
	return fetcher(`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}`, {
		body: JSON.stringify(data),
		method: "PATCH",
	}).then((response) => {
		mutate(`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s`);
		mutate(`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}`);
		return response;
	});
};

/**
 * delete{{title parent}} deletes an existing {{toLower parent}}.
 */
export const delete{{title parent}} = (params, fetcher) => {
	const { {{toLower project}}, {{toLower parent}} } = params;
	return fetcher(`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}`, {
		method: "DELETE",
	}).then((response) => {
		mutate(`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s`);
		return response;
	});
};

/**
 * use returns an swr hook that provides
 */
export const use{{title parent}}List = ({{toLower project}}) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s`
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
export const use{{title parent}} = ({{toLower project}}, {{toLower parent}}) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}`
	);

	return {
		{{toLower parent}}: data,
		isLoading: !error && !data,
		isError: error,
	};
};
