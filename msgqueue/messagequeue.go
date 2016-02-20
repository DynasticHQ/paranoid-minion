/* Paranoid Minion
 * Copyright (C) 2016 Miguel Moll
 *
 * This software is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This software is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this software.  If not, see <http://www.gnu.org/licenses/>.
 */

package msgqueue

// MessageQueue is the interface expected for a message queue system.
type MessageQueue interface {
	Initialize(host, queueName string) error
	Connect() error
	SetIncomingChannel(rc chan<- QueueData)
	Shutdown()
}

type QueueData map[string]interface{}

type MessageConverter interface {
	ToQueueData(b []byte, qd *QueueData) error
}
