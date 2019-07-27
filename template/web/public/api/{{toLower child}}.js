import { instance } from "./config.js";
import useSWR, { mutate } from "swr";

/**
 * create{{title child}} creates a new {{toLower child}}.
 */
export const create{{title child}} = async ({{toLower project}}, {{toLower parent}}, data, fetcher) => {
	return fetcher(
		`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s`,
		{
			body: JSON.stringify(data),
			method: "POST",
		}
	).then(({{toLower child}}) => {
		mutate(`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s`);
		return {{toLower child}};
	});
};

/**
 * update{{title child}} updates an existing {{toLower child}}.
 */
export const update{{title child}} = ({{toLower project}}, {{toLower parent}}, {{toLower child}}, data, fetcher) => {
	return fetcher(
		`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s/${{`{`}}{{toLower child}}{{`}`}}`,
		{
			body: JSON.stringify(data),
			method: "PATCH",
		}
	);
};

/**
 * delete{{title child}} deletes an existing {{toLower child}}.
 */
export const delete{{title child}} = ({{toLower project}}, {{toLower parent}}, {{toLower child}}, fetcher) => {
	return fetcher(
		`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s/${{`{`}}{{toLower child}}{{`}`}}`,
		{
			method: "DELETE",
		}
	).then((_) => {
		mutate(`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s`);
		return;
	});
};

/**
 * use returns an swr hook that provides
 */
export const use{{title child}}List = ({{toLower project}}, {{toLower parent}}) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s`
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
export const use{{title child}} = ({{toLower project}}, {{toLower parent}}, {{toLower child}}) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/{{toLower parent}}s/${{`{`}}{{toLower parent}}{{`}`}}/{{toLower child}}s/${{`{`}}{{toLower child}}{{`}`}}`
	);

	return {
		{{toLower child}}: data,
		isLoading: !error && !data,
		isError: error,
	};
};
