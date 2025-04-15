# PortCtl 🚀

## A lightweight CLI tool to find and kill processes by port or PID. Built with Go and Cobra.

## 🔧 Installation

```
git clone https://github.com/yourusername/portctl.git
cd portctl
go mod tidy
go build -o portctl
```

## ⚙️ Usage

### 🔍 Find process by port
```
./portctl find --port 3000
```

### 💀 Kill process by port
```
./portctl kill --port 3000
```

### 💀 Kill process by PID
```
./portctl kill --pid 12345
```

### 📌 Example
```
./portctl find --port 3000
🔍 Process found: PID 12345

./portctl kill --pid 12345
✅ Process killed!
```