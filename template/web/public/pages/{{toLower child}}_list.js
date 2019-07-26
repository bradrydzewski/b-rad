import { useState, useRef } from "react";
import styles from "./{{toLower child}}_list.module.css";
import { Link } from "wouter";
import { use{{parent}} } from "../api/{{toLower parent}}.js";
import { use{{child}}List, create{{child}} } from "../api/{{toLower child}}.js";
import { useProject } from "../api/project.js";
import { useSession } from "../hooks/session.js";

// Renders the {{child}} List page.
export default function {{child}}List({ params }) {
	const { fetcher } = useSession();

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
	// Load {{child}} List
	//

	const {
		{{toLower child}}List,
		isLoading: is{{child}}Loading,
		isError: is{{child}}Error,
	} = use{{child}}List(params.project, params.{{toLower parent}});

	if (is{{child}}Loading) {
		return renderLoading();
	}
	if (is{{child}}Error) {
		return renderError(is{{child}}Error);
	}

	//
	// Add {{child}} Functions
	//

	const [error, setError] = useState(null);
	const nameElem = useRef(null);
	const descElem = useRef(null);

	const handleCreate = () => {
		const name = nameElem.current.value;
		const desc = descElem.current.value;
		create{{child}}(project.id, {{toLower parent}}.id, { name, desc }, fetcher).then(({{toLower child}}) => {
			nameElem.current.value = "";
			descElem.current.value = "";
		});
	};

	//
	// Render Page
	//

	return (
		<>
			<section className={styles.root}>
				<ul>
					{{`{`}}{{toLower child}}List.map(({{toLower child}}) => (
						<{{child}}Info {{toLower parent}}={{`{`}}{{toLower parent}}{{`}`}} {{toLower child}}={{`{`}}{{toLower child}}{{`}`}} project={project} />
					))}
				</ul>

				<div className="actions">
					<button onClick={handleCreate}>Add {{child}}</button>
					<input ref={nameElem} type="text" placeholder="name" />
					<input ref={descElem} type="text" placeholder="desc" />
				</div>
			</section>
		</>
	);
}

// render the {{toLower child}} information.
const {{child}}Info = ({ {{toLower parent}}, {{toLower child}}, project }) => {
	return (
		<li id={{`{`}}{{toLower parent}}.id}>
			<Link
				href={`/projects/${project.id}/{{toLower parent}}s/${{`{`}}{{toLower parent}}.id}/{{toLower child}}s/${{`{`}}{{toLower child}}.id}`}
			>
				{{`{`}}{{toLower child}}.name}
			</Link>
		</li>
	);
};

// helper function renders the loading bar.
const renderLoading = () => {
	return <div>Loading ...</div>;
};

// helper function returns the error message.
const renderError = (error) => {
	return <div>{error}</div>;
};
