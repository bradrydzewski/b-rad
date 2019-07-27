import styles from "./{{toLower parent}}.module.css";
import { Route, Switch } from "wouter";
import { use{{title parent}} } from "../../api/{{toLower parent}}.js";
import { use{{title project}} } from "../../api/{{toLower project}}.js";

import {{title child}}List from "../{{toLower child}}s/{{toLower child}}s.js";
import {{title child}}Info from "../{{toLower child}}/{{toLower child}}.js";

// Renders the {{title parent}} page.
export default function {{title parent}}({ params }) {
	//
	// Load {{title project}}
	//

	const {
		{{toLower project}},
		isLoading: is{{title project}}Loading,
		isError: is{{title project}}Error,
	} = use{{title project}}(params.{{toLower project}});

	if (is{{title project}}Loading) {
		return renderLoading();
	}
	if (is{{title project}}Error) {
		return renderError(is{{title project}}Error);
	}

	//
	// Load {{title parent}}
	//

	const { {{toLower parent}}, isLoading: is{{title parent}}Loading, isError: is{{title parent}}Errror } = use{{title parent}}(
		params.{{toLower project}},
		params.{{toLower parent}}
	);

	if (is{{title parent}}Loading) {
		return renderLoading();
	}
	if (is{{title parent}}Errror) {
		return renderError(is{{title parent}}Errror);
	}

	//
	// Render Page
	//

	return (
		<>
			<Switch>
				<Route
					path="/{{toLower project}}s/:{{toLower project}}/{{toLower parent}}s/:{{toLower parent}}/{{toLower child}}s/:{{toLower child}}"
					component={{`{`}}{{title child}}Info}
				/>
				<Route
					path="/{{toLower project}}s/:{{toLower project}}/{{toLower parent}}s/:{{toLower parent}}"
					component={{`{`}}{{title child}}List}
				/>
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
