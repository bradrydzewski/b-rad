import { useState, useRef } from "react";
import styles from "./{{toLower project}}s.module.css";
import { Link } from "wouter";
import {
	use{{title project}}List,
	create{{title project}},
	delete{{title project}},
} from "../../api/{{toLower project}}.js";
import { useSession } from "../../hooks/session.js";

import Button from "../../shared/components/button";
import Input from "../../shared/components/input";
import Avatar from "../../shared/components/avatar";

import { Drawer, Target } from "@accessible/drawer";

// Renders the Home page.
export default function Home() {
	const { fetcher } = useSession();
	const [open, setOpen] = useState(false);

	//
	// Load {{title project}} List
	//

	const { {{toLower project}}List, isLoading, isError } = use{{title project}}List();
	if (isLoading) {
		return renderLoading();
	}
	if (isError) {
		return renderError(isError);
	}

	//
	// Create {{title project}} Function
	//

	const [error, setError] = useState(null);
	const nameElem = useRef(null);
	const descElem = useRef(null);

	const handleCreate = () => {
		const name = nameElem.current.value;
		const desc = descElem.current.value;
		create{{title project}}({ name, desc }, fetcher).then(({{toLower project}}) => {
			nameElem.current.value = "";
			descElem.current.value = "";
			setOpen(false);
		});
	};

	//
	// Handle Deletions
	//

	const handleDelete = ({{toLower project}}) => {
		delete{{title project}}({{toLower project}}, fetcher);
	};

	return (
		<>
			<section className={styles.root}>
				<ul className={styles.list}>
					{{`{`}}{{toLower project}}List.map(({{toLower project}}) => (
						<{{title project}}Info {{toLower project}}={{`{`}}{{toLower project}}{{`}`}} onDelete={handleDelete} />
					))}
				</ul>

				<Button className={styles.button} onClick={() => setOpen(true)}>
					New {{title project}}
				</Button>

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
								<Button onClick={handleCreate}>Create {{title project}}</Button>
								<Button onClick={() => setOpen(false)}>Close</Button>
							</div>
						</div>
					</Target>
				</Drawer>
			</section>
		</>
	);
}

// render the {{toLower project}} information.
const {{title project}}Info = ({ {{toLower project}}, onDelete }) => {
	return (
		// <li >
		<Link
			href={`/{{toLower project}}s/${{`{`}}{{toLower project}}.id}`}
			id={{`{`}}{{toLower project}}.id}
			className={styles.item}
		>
			<Avatar text={{`{`}}{{toLower project}}.name} />
			{{`{`}}{{toLower project}}.name}
		</Link>
		///* <button onClick={onDelete.bind(this, {{toLower project}})}>Delete</button> */}
		///* </li> */}
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
	return <div>Your {{title project}} list is empty</div>;
};
