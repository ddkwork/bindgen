
#pragma once

static_assert(sizeof(DEBUGGEE_UD_PAUSED_PACKET) < PacketChunkSize,
              "err (static_assert), size of PacketChunkSize should be bigger "
              "than DEBUGGEE_UD_PAUSED_PACKET");

static_assert(sizeof(DEBUGGER_UPDATE_SYMBOL_TABLE) < PacketChunkSize,
              "err (static_assert), size of PacketChunkSize should be bigger "
              "than DEBUGGER_UPDATE_SYMBOL_TABLE (MODULE_SYMBOL_DETAIL)");
