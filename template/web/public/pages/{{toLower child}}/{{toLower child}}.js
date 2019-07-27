import { useState } from "react";
import styles from "./{{toLower child}}.module.css";
import { useSession } from "../../hooks/session.js";
import { use{{title project}} } from "../../api/{{toLower project}}.js";
import { use{{title parent}} } from "../../api/{{toLower parent}}.js";
import { use{{title child}} } from "../../api/{{toLower child}}.js";
import { Link } from "wouter";

import Button from "../../shared/components/button";
import Breadcrumb from "../../shared/components/breadcrumb";
import Input from "../../shared/components/input";

// Renders the {{title child}} Info page.
export default function {{title child}}({ params }) {
	const { session, fetcher } = useSession();
	const [showToken, setShowToken] = useState(false);

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
		{{toLower child}},
		isLoading: is{{title child}}Loading,
		isError: is{{title child}}Error,
	} = use{{title child}}(params.{{toLower project}}, params.{{toLower parent}}, params.{{toLower child}});

	if (is{{title child}}Loading) {
		return renderLoading();
	}
	if (is{{title child}}Error) {
		return renderError(is{{title child}}Error);
	}

	return (
		<>
			<section className={styles.root}>
				<Breadcrumb>
					<Link href={`/{{toLower project}}s/${{`{`}}{{toLower project}}.id}/{{toLower parent}}s`}>{{title parent}}s</Link>
					<Link href={`/{{toLower project}}s/${{`{`}}{{toLower project}}.id}/{{toLower parent}}s/${{`{`}}{{toLower parent}} && {{toLower parent}}.id}`}>
						{{`{`}}{{toLower parent}} && {{toLower parent}}.name}
					</Link>
					<Link href={`/{{toLower project}}s/${{`{`}}{{toLower project}}.id}/{{toLower parent}}s/${{`{`}}{{toLower parent}} && {{toLower parent}}.id}`}>
						{{title child}}s
					</Link>
				</Breadcrumb>
				<h1>{{title child}}</h1>
				<div className={styles.card}>
					<h2>{{title child}}</h2>
					<div className={styles.field}>
						<label>Name *</label>
						<Input type="text" value={{`{`}}{{toLower child}} && {{toLower child}}.name} />
					</div>
					<div className={styles.field}>
						<label>Description *</label>
						<Input type="text" value={{`{`}}{{toLower child}} && {{toLower child}}.desc} />
					</div>
					<div className={styles.actions}>
						<Button>Update</Button>
					</div>
				</div>

				<div className={styles.card}>
					<h2>Delete {{title child}}</h2>
					<p>Warning, this action cannot be undone.</p>
					<Button>Delete</Button>
				</div>
			</section>
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
