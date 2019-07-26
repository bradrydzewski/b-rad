import styles from "./project.module.css";
import { Route, Switch } from "wouter";
import { use{{parent}} } from "../api/{{toLower parent}}.js";
import { useProject } from "../api/project.js";

import Link from "../shared/link.js";

import {{child}}List from "./{{toLower child}}_list.js";

// Renders the {{parent}} page.
export default function {{parent}}({ params }) {
	//
	// Load Project
	//

	const {
		project,
		isLoading: isProjectLoading,
		isError: isProjectError,
	} = useProject(params.project);

	if (isProjectLoading) {
		return renderLoading();
	}
	if (isProjectError) {
		return renderError(isProjectError);
	}

	//
	// Load {{parent}}
	//

	const { {{toLower parent}}, isLoading: is{{parent}}Loading, isError: is{{parent}}Errror } = use{{parent}}(
		params.project,
		params.{{toLower parent}}
	);

	if (is{{parent}}Loading) {
		return renderLoading();
	}
	if (is{{parent}}Errror) {
		return renderError(is{{parent}}Errror);
	}

	//
	// Render Page
	//

	return (
		<>
			<nav>
				<h1>{{`{`}}{{toLower parent}}.name}</h1>
				<ul>
					<li>
						<Link href={`/projects/${project.id}/{{toLower parent}}s/${{`{`}}{{toLower parent}}.id}`}>
							{{child}}s
						</Link>
					</li>
					<li>
						<Link href="#">Details</Link>
					</li>
				</ul>
			</nav>

			<Switch>
				<Route path="/projects/:project/{{toLower parent}}s/:{{toLower parent}}" component={{`{`}}{{child}}List} />
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
