# Logging Improvements Plan

## Current State
The neural network training process currently outputs too much information to the terminal, making it difficult to track progress and potentially impacting performance. This document outlines a plan to improve the logging system.

## Terminal Output Modes

### 1. Training Mode (Default)
Minimal output focused on essential progress information:
```
[===========>     ] 45.0% Game 450/1000
X: 35% O: 30% Draw: 35% ε=0.1
```

### 2. Verbose Mode
Enabled via `-v` or `--verbose` flag. Shows:
- Current board state
- Move probabilities
- Strategy information (forks, blocks, etc.)
- Detailed statistics

### 3. Statistics Mode
Automatically displays detailed statistics every N games (e.g., every 100):
```
=== Training Statistics (Game 100) ===
Win Rates:
  X: 35.0% (35 wins)
  O: 30.0% (30 wins)
  Draws: 35.0% (35 draws)

Strategic Moves:
  Fork Creations: 25 (25.0% of games)
  Fork Blocks: 20 (20.0% of games)
  Winning Moves: 40 (40.0% of games)
  Blocking Moves: 35 (35.0% of games)

Training Parameters:
  Epsilon: 0.100
  Time since last save: 5m30s
==========================================
```
- Clears screen before showing stats
- Returns to progress bar after 5 seconds

## Log File Improvements

### 1. Structure
- Timestamp for each entry
- Clear section headers
- Consistent formatting

### 2. Content
- Game state transitions
- Move selections and probabilities
- Strategic analysis
- Network weight updates
- Performance metrics

### 3. Organization
- Separate log files for different aspects:
  - `training.log`: General training progress
  - `strategies.log`: Strategic move analysis
  - `network.log`: Network weight updates
  - `performance.log`: Performance metrics

## Implementation Plan

1. **Phase 1: Terminal Output**
   - Implement minimal progress display
   - Add verbose mode flag
   - Create statistics display function

2. **Phase 2: Log File Structure**
   - Implement timestamped logging
   - Create separate log files
   - Add log rotation

3. **Phase 3: Content Organization**
   - Implement strategic analysis logging
   - Add network weight logging
   - Create performance metrics logging

4. **Phase 4: Performance Optimization**
   - Buffer log writes
   - Implement log level filtering
   - Add log compression

## Benefits

1. **Improved Readability**
   - Clear, focused terminal output
   - Well-organized log files
   - Easy access to detailed information

2. **Better Performance**
   - Reduced terminal I/O
   - Buffered log writing
   - Efficient log rotation

3. **Enhanced Debugging**
   - Comprehensive log history
   - Structured log format
   - Separate concerns in different log files

## Next Steps

1. Create new logging package
2. Implement progress bar display
3. Add verbose mode flag
4. Set up log file structure
5. Implement statistics display
6. Add log rotation
7. Create performance metrics
8. Add network weight logging
9. Implement strategic analysis logging
10. Add log compression 