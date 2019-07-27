import { useState, useEffect } from "react";
import { useLocation } from "wouter";
import styles from "./settings.module.css";
import { useSession } from "../../hooks/session.js";
import { use{{title project}} } from "../../api/{{toLower project}}.js";
import { use{{title parent}}, update{{title parent}}, delete{{title parent}} } from "../../api/{{toLower parent}}.js";

import Button from "../../shared/components/button";
import Input from "../../shared/components/input";

// Renders the {{title project}} Settings page.
export default function Settings({ params }) {
	const { fetcher } = useSession();
	const [_, setLocation] = useLocation();

	//
	// Load {{title project}}
	//

	const { {{toLower project}} } = use{{title project}}(params.{{toLower project}});

	//
	// Load {{title parent}}
	//

	const { {{toLower parent}} } = use{{title parent}}(params.{{toLower project}}, params.{{toLower parent}});
	const [{{toLower parent}}Data, set{{title parent}}Data] = useState({});
	useEffect(() => {{toLower parent}} && set{{title parent}}Data({{toLower parent}}), [{{toLower parent}}]);

	//
	// Update {{title parent}}
	//

	const handleUpdateName = (event) => {
		set{{title parent}}Data({
			name: event.target.value,
			desc: {{toLower parent}}Data.desc,
		});
	};

	const handleUpdateDesc = (event) => {
		set{{title parent}}Data({
			desc: event.target.value,
			name: {{toLower parent}}Data.name,
		});
	};

	const handleUpdate = () => {
		const params = {
			{{toLower project}}: {{toLower project}}.id,
			{{toLower parent}}: {{toLower parent}}.id,
		};
		update{{title parent}}(params, {{toLower parent}}Data, fetcher);
	};

	//
	// Delete {{title parent}}
	//

	const handleDelete = () => {
		if (confirm("Are you sure you want to proceed?")) {
			const params = {
				{{toLower project}}: {{toLower project}}.id,
				{{toLower parent}}: {{toLower parent}}.id,
			};
			delete{{title parent}}(params, fetcher);
			setLocation(`/{{toLower project}}s/${{`{`}}{{toLower project}}.id}`);
		}
	};

	return (
		<>
			<section className={styles.root}>
				<div className={styles.card}>
					<h2>{{title parent}}</h2>
					<div className={styles.field}>
						<label>Name *</label>
						<Input
							type="text"
							value={{`{`}}{{toLower parent}}Data.name}
							onChange={handleUpdateName}
						/>
					</div>
					<div className={styles.field}>
						<label>Description *</label>
						<Input
							type="text"
							value={{`{`}}{{toLower parent}}Data.desc}
							onChange={handleUpdateDesc}
						/>
					</div>
					<div className={styles.actions}>
						<Button onClick={handleUpdate}>Update {{title project}}</Button>
					</div>
				</div>

				<div className={styles.card}>
					<h2>Delete</h2>
					<p>Warning, this action cannot be undone.</p>
					<Button onClick={handleDelete}>Delete</Button>
				</div>
			</section>
		</>
	);
}
