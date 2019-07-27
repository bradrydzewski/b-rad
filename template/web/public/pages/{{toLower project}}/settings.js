import { useState, useEffect } from "react";
import { useLocation } from "wouter";
import styles from "./settings.module.css";
import { useSession } from "../../hooks/session.js";
import { use{{title project}}, update{{title project}}, delete{{title project}} } from "../../api/{{toLower project}}.js";

import Button from "../../shared/components/button";
import Input from "../../shared/components/input";

// Renders the {{title project}} Settings page.
export default function Settings({ params }) {
	const { fetcher } = useSession();
	const [showToken, setShowToken] = useState(false);
	const [_, setLocation] = useLocation();

	//
	// Load {{title project}}
	//

	const { {{toLower project}} } = use{{title project}}(params.{{toLower project}});
	const [{{toLower project}}Data, set{{title project}}Data] = useState({});
	useEffect(() => {{toLower project}} && set{{title project}}Data({{toLower project}}), [{{toLower project}}]);

	//
	// Update {{title project}}
	//

	const handleUpdateName = (event) => {
		set{{title project}}Data({
			name: event.target.value,
			desc: {{toLower project}}Data.desc,
		});
	};

	const handleUpdateDesc = (event) => {
		set{{title project}}Data({
			desc: event.target.value,
			name: {{toLower project}}Data.name,
		});
	};

	const handleUpdate = () => {
		update{{title project}}({{toLower project}}, {{toLower project}}Data, fetcher);
	};

	//
	// Delete {{title project}}
	//

	const handleDelete = () => {
		if (confirm("Are you sure you want to proceed?")) {
			delete{{title project}}({{toLower project}}, fetcher);
			setLocation("/");
		}
	};

	return (
		<>
			<section className={styles.root}>
				<div className={styles.card}>
					<h2>{{title project}}</h2>
					<div className={styles.field}>
						<label>Name *</label>
						<Input
							type="text"
							value={{`{`}}{{toLower project}}Data.name}
							onChange={handleUpdateName}
						/>
					</div>
					<div className={styles.field}>
						<label>Description *</label>
						<Input
							type="text"
							value={{`{`}}{{toLower project}}Data.desc}
							onChange={handleUpdateDesc}
						/>
					</div>
					<div className={styles.actions}>
						<Button onClick={handleUpdate}>Update {{title project}}</Button>
					</div>
				</div>

				<div className={styles.card}>
					<h2>Token</h2>
					<p>{{title project}} access token that can be used to access the API.</p>
					{showToken && <pre>{{`{`}}{{toLower project}} && {{toLower project}}.token}</pre>}
					{!showToken && (
						<Button onClick={() => setShowToken(true)}>Display Token</Button>
					)}
				</div>

				<div className={styles.card}>
					<h2>Delete {{title project}}</h2>
					<p>Warning, this action cannot be undone.</p>
					<Button onClick={handleDelete}>Delete</Button>
				</div>
			</section>
		</>
	);
}
