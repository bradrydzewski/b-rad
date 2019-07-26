import { instance } from "./config.js";
import useSWR, { mutate } from "swr";

/**
 * create{{child}} creates a new {{toLower child}}.
 */
export const create{{child}} = async (project, {{toLower parent}}, data, fetcher) => {
	return fetcher(
		`${instance}/api/v1/projects/${project}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s`,
		{
			body: JSON.stringify(data),
			method: "POST",
		}
	).then(({{toLower child}}) => {
		mutate(`${instance}/api/v1/projects/${project}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s`);
		return {{toLower child}};
	});
};

/**
 * update{{child}} updates an existing {{toLower child}}.
 */
export const update{{child}} = (project, {{toLower parent}}, {{toLower child}}, data, fetcher) => {
	return fetcher(
		`${instance}/api/v1/projects/${project}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s/${{`{`}}{{toLower child}}{{`}`}}`,
		{
			body: JSON.stringify(data),
			method: "PATCH",
		}
	);
};

/**
 * delete{{child}} deletes an existing {{toLower child}}.
 */
export const delete{{child}} = (project, {{toLower parent}}, {{toLower child}}, fetcher) => {
	return fetcher(
		`${instance}/api/v1/projects/${project}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s/${{`{`}}{{toLower child}}{{`}`}}`,
		{
			method: "DELETE",
		}
	);
};

/**
 * use returns an swr hook that provides
 */
export const use{{child}}List = (project, {{toLower parent}}) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/projects/${project}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s`
	);

	return {
		{{toLower child}}List: data,
		isLoading: !error && !data,
		isError: error,
	};
};

/**
 * use returns an swr hook that provides
 */
export const use{{child}} = (project, {{toLower parent}}, {{toLower child}}) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/projects/${project}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s/${{`{`}}{{toLower child}}{{`}`}}`
	);

	return {
		{{toLower child}}: data,
		isLoading: !error && !data,
		isError: error,
	};
};
