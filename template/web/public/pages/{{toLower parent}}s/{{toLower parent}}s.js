import { useState, useRef } from "react";
import styles from "./{{toLower parent}}s.module.css";
import { Link } from "wouter";
import { use{{title parent}}List, create{{title parent}}, delete{{title parent}} } from "../../api/{{toLower parent}}.js";
import { use{{title project}} } from "../../api/{{toLower project}}.js";
import { useSession } from "../../hooks/session.js";

import Avatar from "../../shared/components/avatar";
import Button from "../../shared/components/button";
import Input from "../../shared/components/input";

import { Drawer, Target } from "@accessible/drawer";

// Renders the {{title parent}} List page.
export default function {{title parent}}List({ params }) {
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
	// Load {{title parent}} List
	//

	const {
		{{toLower parent}}List,
		isLoading: is{{title parent}}Loading,
		isError: is{{title parent}}Errror,
	} = use{{title parent}}List({{toLower project}} && {{toLower project}}.id);

	if (is{{title parent}}Loading) {
		return renderLoading();
	}
	if (is{{title parent}}Errror) {
		return renderError(is{{title parent}}Errror);
	}

	//
	// Add {{title parent}} Functions
	//

	const [error, setError] = useState(null);
	const nameElem = useRef(null);
	const descElem = useRef(null);

	const handleCreate = () => {
		const name = nameElem.current.value;
		const desc = descElem.current.value;
		const data = { name, desc };
		const params = { {{toLower project}}: {{toLower project}}.id };
		create{{title parent}}(params, data, fetcher).then(({{toLower project}}) => {
			nameElem.current.value = "";
			descElem.current.value = "";
			setOpen(false);
		});
	};

	//
	// Handle Deletions
	//

	const handleDelete = ({{toLower parent}}) => {
		const params = { {{toLower project}}: {{toLower project}}.id, {{toLower parent}}: {{toLower parent}}.id };
		delete{{title parent}}(params, fetcher);
	};

	//
	// Render Page
	//

	return (
		<>
			<section className={styles.root}>
				<ul className={styles.list}>
					{{`{`}}{{toLower parent}}List.map(({{toLower parent}}) => (
						<{{title parent}}Info
							{{toLower parent}}={{`{`}}{{toLower parent}}{{`}`}}
							{{toLower project}}={{`{`}}{{toLower project}}{{`}`}}
							onDelete={handleDelete}
						/>
					))}
				</ul>

				<Button className={styles.button} onClick={() => setOpen(true)}>
					New {{title parent}}
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
							<Button onClick={handleCreate}>Add {{title parent}}</Button>
							<Button onClick={() => setOpen(false)}>Close</Button>
						</div>
					</div>
				</Target>
			</Drawer>
		</>
	);
}

// render the {{toLower parent}} information.
const {{title parent}}Info = ({ {{toLower parent}}, {{toLower project}}, onDelete }) => {
	return (
		<li id={{`{`}}{{toLower parent}}.id} className={styles.item}>
			<Avatar text={{`{`}}{{toLower parent}}.name} className={styles.avatar} />
			<Link
				href={`/{{toLower project}}s/${{`{`}}{{toLower project}}.id}/{{toLower parent}}s/${{`{`}}{{toLower parent}}.id}`}
				className={styles.fill}
			>
				{{`{`}}{{toLower parent}}.name}
			</Link>
			<Button onClick={onDelete.bind(this, {{toLower parent}})}>Delete</Button>
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
	return <div>Your {{title parent}} list is empty</div>;
};
