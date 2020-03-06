/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package administration

type ClusterBackupInfo struct {

	// IP address of the node from which the backup was taken
	IpAddress string `json:"ip_address,omitempty"`

	// ID of the node from which the backup was taken
	NodeId string `json:"node_id,omitempty"`

	// timestamp of the cluster backup file
	Timestamp int64 `json:"timestamp,omitempty"`
}