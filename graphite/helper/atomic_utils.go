// replication-manager - Replication Manager Monitoring and CLI for MariaDB and MySQL
// Copyright 2017 Signal 18 SARL
// Authors: Guillaume Lefranc <guillaume@signal18.io>
//          Stephane Varoqui  <svaroqui@gmail.com>
// This source code is licensed under the GNU General Public License, version 3.

package helper

import "sync/atomic"

type StatCallback func(metric string, value float64)

func SendAndSubstractUint64(metric string, v *uint64, send StatCallback) {
	res := atomic.LoadUint64(v)
	atomic.AddUint64(v, ^uint64(res-1))
	send(metric, float64(res))
}

func SendUint64(metric string, v *uint64, send StatCallback) {
	res := atomic.LoadUint64(v)
	send(metric, float64(res))
}

func SendUint32(metric string, v *uint32, send StatCallback) {
	res := atomic.LoadUint32(v)
	send(metric, float64(res))
}

func SendAndSubstractUint32(metric string, v *uint32, send StatCallback) {
	res := atomic.LoadUint32(v)
	atomic.AddUint32(v, ^uint32(res-1))
	send(metric, float64(res))
}

func SendAndZeroIfNotUpdatedUint32(metric string, v *uint32, send StatCallback) {
	res := atomic.LoadUint32(v)
	atomic.CompareAndSwapUint32(v, res, 0)
	send(metric, float64(res))
}
