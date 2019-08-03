import { useState, useRef } from "react";
import styles from "./members.module.css";
import { useMemberList, createMember, deleteMember } from "../../api/member.js";
import { use{{title project}} } from "../../api/{{toLower project}}.js";
import { useSession } from "../../hooks/session.js";

import Avatar from "../../shared/components/avatar";
import Button from "../../shared/components/button";
import Input from "../../shared/components/input";
import Select from "../../shared/components/select";

import { Drawer, Target } from "@accessible/drawer";

// Renders the Member List page.
export default function Members({ params }) {
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
	// Load Member List
	//

	const {
		memberList,
		isLoading: isMemberLoading,
		isError: isMemberErrror,
	} = useMemberList({{toLower project}} && {{toLower project}}.slug);

	if (isMemberLoading) {
		return renderLoading();
	}
	if (isMemberErrror) {
		return renderError(isMemberErrror);
	}
	if (memberList.length === 0) {
		return renderEmpty();
	}

	//
	// Add Member Functions
	//

	const [error, setError] = useState(null);
	const emailElem = useRef(null);
	const roleElem = useRef(null);

	const handleCreate = () => {
		const email = emailElem.current.value;
		const role = roleElem.current.value;
		createMember({{toLower project}}.slug, { email, role }, fetcher).then((member) => {
			emailElem.current.value = "";
			setOpen(false);
		});
	};

	//
	// Delete Functions
	//

	const handleDelete = ({{toLower project}}, user) => {
		if (confirm("Are you sure you want to proceed?")) {
			deleteMember({{toLower project}}.slug, user.email, fetcher);
		}
	};

	//
	// Render Page
	//

	return (
		<>
			<section className={styles.root}>
				<ul className={styles.list}>
					{memberList.map((member) => (
						<MemberInfo
							member={member}
							{{toLower project}}={{`{`}}{{toLower project}}{{`}`}}
							onDelete={handleDelete}
						/>
					))}
				</ul>

				<Button className={styles.button} onClick={() => setOpen(true)}>
					New Member
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
						<Input ref={emailElem} type="text" placeholder="email" />
						<Select ref={roleElem}>
							<option value="developer" selected>
								developer
							</option>
							<option value="admin">admin</option>
						</Select>

						<div className={styles.actions}>
							<Button onClick={handleCreate}>Add Member</Button>
							<Button onClick={() => setOpen(false)}>Close</Button>
						</div>
					</div>
				</Target>
			</Drawer>
		</>
	);
}

// render the member information.
const MemberInfo = ({ member, {{toLower project}}, onDelete }) => {
	return (
		<li id={member.slug} className={styles.item}>
			<Avatar text={member.email} className={styles.avatar} />
			<span className={styles.fill}>
				{member.email} ({member.role})
			</span>
			<Button onClick={onDelete.bind(this, {{toLower project}}, member)}>Delete</Button>
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
	return <div>Your Member list is empty</div>;
};
