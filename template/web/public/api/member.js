import { instance } from "./config.js";
import useSWR, { mutate } from "swr";

/**
 * createMember creates a new member.
 */
export const createMember = ({{toLower project}}, data, fetcher) => {
	return fetcher(
		`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/members/${data.email}`,
		{
			body: JSON.stringify(data),
			method: "POST",
		}
	).then((member) => {
		mutate(`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/members`);
		return member;
	});
};

/**
 * updateMember updates an existing member.
 */
export const updateMember = ({{toLower project}}, member, data, fetcher) => {
	return fetcher(`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/members/${member}`, {
		body: JSON.stringify(data),
		method: "PATCH",
	});
};

/**
 * deleteMember deletes an existing member.
 */
export const deleteMember = ({{toLower project}}, member, fetcher) => {
	return fetcher(`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/members/${member}`, {
		method: "DELETE",
	}).then((response) => {
		mutate(`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/members`);
		return response;
	});
};

/**
 * use returns an swr hook that provides
 */
export const useMemberList = ({{toLower project}}) => {
	const { data, error } = useSWR(
		`${instance}/api/v1/{{toLower project}}s/${{`{`}}{{toLower project}}{{`}`}}/members`
	);

	return {
		memberList: data,
		isLoading: !error && !data,
		isError: error,
	};
};
