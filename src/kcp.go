/*
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * Author: FTwOoO <booobooob@gmail.com>
 */

package hop

import "net"

type KCPListener struct {

}


// Accept waits for and returns the next connection to the listener.
func (conn *KCPListener)Accept() (net.Conn, error)

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (conn *KCPListener) Close() error

// Addr returns the listener's network address.
func (conn *KCPListener) Addr() net.Addr
