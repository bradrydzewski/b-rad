import { useState, useRef } from "react";
import styles from "./{{toLower parent}}_list.module.css";
import { Link } from "wouter";
import { use{{parent}}List, create{{parent}}, delete{{parent}} } from "../api/{{toLower parent}}.js";
import { useProject } from "../api/project.js";
import { useSession } from "../hooks/session.js";

// Renders the {{parent}} List page.
export default function {{parent}}List({ params }) {
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
	// Load {{parent}} List
	//

	const {
		{{toLower parent}}List,
		isLoading: is{{parent}}Loading,
		isError: is{{parent}}Errror,
	} = use{{parent}}List(project && project.id);

	if (is{{parent}}Loading) {
		return renderLoading();
	}
	if (is{{parent}}Errror) {
		return renderError(is{{parent}}Errror);
	}

	//
	// Add {{parent}} Functions
	//

	const [error, setError] = useState(null);
	const nameElem = useRef(null);
	const descElem = useRef(null);

	const handleCreate = () => {
		const name = nameElem.current.value;
		const desc = descElem.current.value;
		const data = { name, desc };
		const params = { project: project.id };
		create{{parent}}(params, data, fetcher).then((project) => {
			nameElem.current.value = "";
			descElem.current.value = "";
		});
	};

	//
	// Handle Deletions
	//

	const handleDelete = ({{toLower parent}}) => {
		const params = { project: project.id, {{toLower parent}}: {{toLower parent}}.id };
		delete{{parent}}(params, fetcher);
	};

	//
	// Render Page
	//

	return (
		<>
			<section className={styles.root}>
				<ul>
					{{`{`}}{{toLower parent}}List.map(({{toLower parent}}) => (
						<{{parent}}Info
							{{toLower parent}}={{`{`}}{{toLower parent}}{{`}`}}
							project={project}
							onDelete={handleDelete}
						/>
					))}
				</ul>

				<div className="actions">
					<button onClick={handleCreate}>Add {{parent}}</button>
					<input ref={nameElem} type="text" placeholder="name" />
					<input ref={descElem} type="text" placeholder="desc" />
				</div>
			</section>
		</>
	);
}

// render the {{toLower parent}} information.
const {{parent}}Info = ({ {{toLower parent}}, project, onDelete }) => {
	return (
		<li id={{`{`}}{{toLower parent}}.id}>
			<Link href={`/projects/${project.id}/{{toLower parent}}s/${{`{`}}{{toLower parent}}.id}`}>
				{{`{`}}{{toLower parent}}.name}
			</Link>
			<button onClick={onDelete.bind(this, {{toLower parent}})}>Delete</button>
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

// helper function returns the empty message.
const renderEmpty = (error) => {
	return <div>Your {{parent}} list is empty</div>;
};
