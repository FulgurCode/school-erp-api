# School erp api

This is school-erp api written in GO.

## Usage

1. Install and setup [Go](https://go.dev/doc/install) and [Mongodb](https://www.mongodb.com/docs/manual/installation/)

2. Clone this repository and change directory
   
   ```shell
   git clone https://github.com/FulgurCode/school-erp-api
   cd school-erp-api
   ```

3. Create `.env` file
   
   ```shell
   touch .env
   ```

4. Store thease values in `.env`
   
   * `PORT`
     Port that server listen for requests
   
   * `SECRET_KEY`
     Secret key for cookie
   
   * `MONGODB_URI`
     Uri of mongodb database
   
   **Example**
   
   ```env
   PORT=9000
   SECRET_KEY=secret_key
   MONGODB_URI=mongodb://localhost:27017
   ```

5. Install go dependencies
   
   ```shell
   go mod tidy
   ```

6. Build and run executable file
   
   ```shell
   go build .
   ```
   
   Now run the executable for your OS

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## LICENSE

[The GPLv3 License (GPLv3)](LICENSE)

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
