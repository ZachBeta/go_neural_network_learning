# Logging Improvements Checklist

## Terminal Output
- [ ] Default display: Simple progress bar with game count
  ```
  [===========>     ] 45.0% Game 450/1000
  ```
- [ ] Statistics display every 100 games
  ```
  === Training Statistics (Game 100) ===
  X: 35% O: 30% Draw: 35% ε=0.1
  ==========================================
  ```
- [ ] Move all other output to log files

## File Logging
- [ ] Create timestamped log entries
- [ ] Log game state transitions
- [ ] Log move selections and probabilities
- [ ] Log strategic analysis (forks, blocks, etc.)
- [ ] Log network weight updates
- [ ] Log performance metrics

## Code Changes
- [ ] Remove all `fmt.Print` statements from training code
- [ ] Create logging package with file writer
- [ ] Implement progress bar display function
- [ ] Implement statistics display function
- [ ] Add log rotation to prevent files from growing too large
- [ ] Add error handling for log file operations

## Files to Modify
- [ ] `cmd/neural_train/main.go`
- [ ] `cmd/neural_train/training.go`
- [ ] `cmd/neural_train/visualization.go`
- [ ] `pkg/neural/utils.go`

## Testing
- [ ] Verify progress bar updates correctly
- [ ] Verify statistics display every 100 games
- [ ] Verify all output is properly logged to files
- [ ] Verify log rotation works correctly
- [ ] Test error handling for log file operations 