import styles from "./project.module.css";
import { Route, Switch } from "wouter";
import { useProject } from "../api/project.js";

import Link from "../shared/link.js";

import {{parent}} from "./{{toLower parent}}.js";
import {{parent}}List from "./{{toLower parent}}_list.js";
import Member from "./members.js";

// Renders the Project page.
export default function Project({ params }) {
	const { project, isLoading, isError } = useProject(params.project);

	if (isLoading) {
		return renderLoading();
	}
	if (isError) {
		return renderError(isError);
	}

	return (
		<>
			<nav>
				<h1>{project && project.name}</h1>
				<ul>
					<li>
						<Link href={`/projects/${project.id}`}>{{parent}}s</Link>
					</li>
					<li>
						<Link href={`/projects/${project.id}/members`}>Members</Link>
					</li>
					<li>
						<Link href={`/projects/${project.id}/settings`}>Details</Link>
					</li>
				</ul>
			</nav>

			<Switch>
				<Route path="/projects/:project" component={{`{`}}{{parent}}List} />
				<Route path="/projects/:project/{{toLower parent}}s" component={{`{`}}{{parent}}List} />
				<Route path="/projects/:project/{{toLower parent}}s/:{{toLower parent}}" component={{`{`}}{{parent}}{{`}`}} />
				<Route path="/projects/:project/{{toLower parent}}s/path+" component={{`{`}}{{parent}}{{`}`}} />
				<Route path="/projects/:project/members" component={Member} />
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
