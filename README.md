# Blackjack CLI

A command-line Blackjack game built for the Linux terminal. This application allows you to create and manage users, play full games of Blackjack, and simulate rounds using various strategies.


## Requirements

* **Go:** 1.24.3


## Installation

Clone the repository, then run:

```
make install
```

After installation:

* A directory named `blackjack` will be created in your home directory.
* The executable will be installed to `/usr/local/bin`, allowing the `blackjack` command to be used system-wide.

## Update
To update the application, run:
```
make update
```

This will:

* Replace the contents of the `blackjack` directory with the udpated game files.


## Uninstallation

To remove the application, run:

```
make uninstall
```

This will:

* Remove the `blackjack` directory from your home folder.
* Remove the executable from `/usr/local/bin`.


## Usage

### User Management

#### 1. Create User

Create a user to store game progress:

```
blackjack user create --username <username> --password <password>
```

#### 2. Sign In

You must sign in before playing:

```
blackjack user signin --username <username> --password <password>
```

#### 3. Sign Out

Sign out when you're done playing:

```
blackjack user signout --username <username> --password <password>
```

#### 4. Retrieve User Statistics

Retrieve statistics for a user:

```
blackjack user retrieve --username <username>
```

**Note:** Use username `HOUSE` to view house statistics.

#### 5. Delete User

Delete a user account:

```
blackjack user delete --username <username> --password <password>
```


## Game Management

### 1. Play Blackjack

Play a standard game against the dealer:

```
blackjack game play
```

**Note:** You must be signed in.

### 2. Simulate Blackjack

Simulate rounds using a chosen strategy:

```
blackjack game simulate --strategy basic --rounds 100
```

Supported strategies:

* `basic`
* `s17`
* `h17`
* `enhc`


## Help

For help with any command:

```
blackjack --help
```

Or:

```
blackjack <category> <command> --help
```

# Enjoy!
