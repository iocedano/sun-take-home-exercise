import React from "react";
import clsx from 'clsx'
import './Status.css'

const OK_STATUS_CODE = 200; 

const Status = ({ url, statusCode }) => {
  return (
    <li className="status">
      <span className="url">{url}</span>
      <i className={clsx('checker', {
        ok: statusCode === OK_STATUS_CODE,
        bad: statusCode !== OK_STATUS_CODE
      })} />
    </li>
  )
}

export default Status;