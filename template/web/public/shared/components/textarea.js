import styles from "./textarea.module.css";
import classnames from "classnames";

export default (props) => (
	<textarea
		className={classnames(styles.root, props.className)}
		ref={props.ref}
		onClick={props.onClick}
		onChange={props.onChange}
		onMouseEnter={props.onMouseEnter}
		onMouseLeave={props.onMouseLeave}
		disabled={props.disabled}
		spellcheck={props.spellCheck}
		placeholder={props.placeholder}
	>
		{props.children}
	</textarea>
);
