import { Route, Switch } from "wouter";
import { use{{title project}} } from "../../api/{{toLower project}}.js";

import {{title parent}} from "../{{toLower parent}}/{{toLower parent}}.js";
import {{title parent}}List from "../{{toLower parent}}s/{{toLower parent}}s.js";
import {{title parent}}Settings from "../{{toLower parent}}/settings.js";
import Member from "../members/members.js";
import Settings from "./settings.js";

// Renders the {{title project}} page.
export default function {{title project}}({ params }) {
	const { {{toLower project}}, isLoading, isError } = use{{title project}}(params.{{toLower project}});

	if (isLoading) {
		return renderLoading();
	}
	if (isError) {
		return renderError(isError);
	}

	return (
		<>
			<Switch>
				<Route path="/{{toLower project}}s/:{{toLower project}}" component={{`{`}}{{title parent}}List} />
				<Route path="/{{toLower project}}s/:{{toLower project}}/{{toLower parent}}s" component={{`{`}}{{title parent}}List} />
				<Route path="/{{toLower project}}s/:{{toLower project}}/{{toLower parent}}s/:{{toLower parent}}" component={{`{`}}{{title parent}}{{`}`}} />
				<Route path="/{{toLower project}}s/:{{toLower project}}/{{toLower parent}}s/:{{toLower parent}}/settings" component={{`{`}}{{title parent}}Settings} />
				<Route path="/{{toLower project}}s/:{{toLower project}}/{{toLower parent}}s/:{{toLower parent}}/{{toLower child}}s/:{{toLower child}}" component={{`{`}}{{title parent}}{{`}`}} />
				<Route path="/{{toLower project}}s/:{{toLower project}}/{{toLower parent}}s/path+" component={{`{`}}{{title parent}}{{`}`}} />
				<Route path="/{{toLower project}}s/:{{toLower project}}/members" component={Member} />
				<Route path="/{{toLower project}}s/:{{toLower project}}/settings" component={Settings} />
				<Route>Not Found</Route>
			</Switch>
		</>
	);
}

// helper function renders the loading bar.
const renderLoading = () => {
	return <div>Loading ...</div>;
};

// helper function returns the error message.
const renderError = (error) => {
	return <div>{error}</div>;
};
