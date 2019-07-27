import { useState, useRef } from "react";
import styles from "./{{toLower child}}s.module.css";
import { Link } from "wouter";
import { use{{title parent}} } from "../../api/{{toLower parent}}.js";
import {
	use{{title child}}List,
	create{{title child}},
	delete{{title child}},
} from "../../api/{{toLower child}}.js";
import { use{{title project}} } from "../../api/{{toLower project}}.js";
import { useSession } from "../../hooks/session.js";

import Avatar from "../../shared/components/avatar";
import Button from "../../shared/components/button";
import Input from "../../shared/components/input";
import Breadcrumb from "../../shared/components/breadcrumb";

import { Drawer, Target } from "@accessible/drawer";

// Renders the {{title child}} List page.
export default function {{title child}}List({ params }) {
	const { fetcher } = useSession();
	const [open, setOpen] = useState(false);

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
	// Load {{title child}} List
	//

	const {
		{{toLower child}}List,
		isLoading: is{{title child}}Loading,
		isError: is{{title child}}Error,
	} = use{{title child}}List(params.{{toLower project}}, params.{{toLower parent}});

	if (is{{title child}}Loading) {
		return renderLoading();
	}
	if (is{{title child}}Error) {
		return renderError(is{{title child}}Error);
	}

	//
	// Add {{title child}} Functions
	//

	const [error, setError] = useState(null);
	const nameElem = useRef(null);
	const descElem = useRef(null);

	const handleCreate = () => {
		const name = nameElem.current.value;
		const desc = descElem.current.value;
		create{{title child}}({{toLower project}}.id, {{toLower parent}}.id, { name, desc }, fetcher).then(
			({{toLower child}}) => {
				nameElem.current.value = "";
				descElem.current.value = "";
				setOpen(false);
			}
		);
	};

	//
	// Delete {{title child}} Functions
	//

	const handleDelete = ({{toLower project}}, {{toLower parent}}, {{toLower child}}) => {
		delete{{title child}}({{toLower project}}.id, {{toLower parent}}.id, {{toLower child}}.id, fetcher);
	};

	//
	// Render Page
	//

	return (
		<>
			<section className={styles.root}>
				<Breadcrumb>
					<Link href={`/{{toLower project}}s/${{`{`}}{{toLower project}}.id}/{{toLower parent}}s`}>{{title parent}}s</Link>
					<Link href={`/{{toLower project}}s/${{`{`}}{{toLower project}}.id}/{{toLower parent}}s/${{`{`}}{{toLower parent}} && {{toLower parent}}.id}`}>
						{{`{`}}{{toLower parent}} && {{toLower parent}}.name}
					</Link>
				</Breadcrumb>

				<h1>{{title child}}s</h1>
				<ul className={styles.list}>
					{{`{`}}{{toLower child}}List.map(({{toLower child}}) => (
						<{{title child}}Info
							{{toLower parent}}={{`{`}}{{toLower parent}}{{`}`}}
							{{toLower child}}={{`{`}}{{toLower child}}{{`}`}}
							{{toLower project}}={{`{`}}{{toLower project}}{{`}`}}
							onDelete={handleDelete}
						/>
					))}
				</ul>

				<Button className={styles.button} onClick={() => setOpen(true)}>
					New {{title child}}
				</Button>
			</section>

			<Drawer open={open}>
				<Target
					placement="right"
					closeOnEscape={true}
					preventScroll={true}
					openClass={styles.drawer}
				>
					<div>
						<Input ref={nameElem} type="text" placeholder="name" />
						<Input ref={descElem} type="text" placeholder="desc" />

						<div className={styles.actions}>
							<Button onClick={handleCreate}>Add {{title child}}</Button>
							<Button onClick={() => setOpen(false)}>Close</Button>
						</div>
					</div>
				</Target>
			</Drawer>
		</>
	);
}

// render the {{toLower child}} information.
const {{title child}}Info = ({ {{toLower parent}}, {{toLower child}}, {{toLower project}}, onDelete }) => {
	return (
		<li id={{`{`}}{{toLower child}}.id} className={styles.item}>
			<Avatar text={{`{`}}{{toLower child}}.name} className={styles.avatar} />
			<Link
				href={`/{{toLower project}}s/${{`{`}}{{toLower project}}.id}/{{toLower parent}}s/${{`{`}}{{toLower parent}}.id}/{{toLower child}}s/${{`{`}}{{toLower child}}.id}`}
				className={styles.fill}
			>
				{{`{`}}{{toLower child}}.name}
			</Link>
			<Button onClick={onDelete.bind(this, {{toLower project}}, {{toLower parent}}, {{toLower child}})}>
				Delete
			</Button>
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
